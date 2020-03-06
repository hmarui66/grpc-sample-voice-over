//
//  ContentViewModel.swift
//  grpcvoiceover
//
//  Created by 丸井弘嗣 on 2020/02/28.
//  Copyright © 2020 hmarui66. All rights reserved.
//

import GRPC
import NIO
import NIOSSL
import Combine
import Foundation
import AVFoundation

final class ContentViewModel: ObservableObject {
    @Published(initialValue: false) var isExecuting: Bool
    @Published(initialValue: Comment_Comment()) var comment: Comment_Comment

    private let speechSynthesizer : AVSpeechSynthesizer

    init() {
        self.speechSynthesizer = AVSpeechSynthesizer()
    }

    func subscribe() {
        guard self.isExecuting == false else {
            print("already started")
            return
        }

        start()
        DispatchQueue.global().async {
            do {
                // See: https://github.com/apple/swift-nio#eventloops-and-eventloopgroups
                let group = MultiThreadedEventLoopGroup(numberOfThreads: 1)

                // Make sure the group is shutdown when we're done with it.
                defer {
                    try! group.syncShutdownGracefully()
                }

                let path = Bundle.main.path(forResource: "ca", ofType: "pem")

                let certificates: [NIOSSLCertificate] = try NIOSSLCertificate.fromPEMFile(path!)

                let tls = ClientConnection.Configuration.TLS(
                    certificateChain: certificates.map { .certificate($0) },
                    certificateVerification: .none,
                    hostnameOverride: "x.test.youtube.com"
                )

                // Provide some basic configuration for the connection, in this case we connect to an endpoint on
                // localhost at the given port.
                let configuration = ClientConnection.Configuration(
                    target: .hostAndPort("10.36.27.189", 8080),
                    eventLoopGroup: group,
                    tls: tls
                )

                // Create a connection using the configuration.
                let connection = ClientConnection(configuration: configuration)

                // Provide the connection to the generated client.
                let client = Comment_CommentServiceServiceClient(channel: connection)

                let req = Comment_Filter.with {
                    $0.query = "test query"
                }

                let call = client.getComment(req) { comment in
                    print(comment)
                    self.setComment(comment)
                }

                let status = try call.status.recover { _ in .processingError }.wait()
                if status.code != .ok {
                    print("PRC failed: \(status)")
                    self.finish()
                }
            } catch {
                print(error)
                self.finish()
            }
        }
    }

    func start() {
        DispatchQueue.main.async {
            self.isExecuting = true
        }
    }

    func finish() {
        DispatchQueue.main.async {
            self.isExecuting = false
        }
    }

    private func setComment(_ comment: Comment_Comment) {
        DispatchQueue.main.async {
            self.comment = comment
        }
        let utterance = AVSpeechUtterance(string: comment.text)
        utterance.voice = AVSpeechSynthesisVoice(language: "ja-JP")
        utterance.rate = 0.7
        utterance.pitchMultiplier = getPitch(comment.user)
        self.speechSynthesizer.speak(utterance)
    }

    private func getPitch(_ user: String) -> Float {
        let num = Array(user.utf8).reduce(0, { x,y in
            x + Float(y)
        })
        let pitch = num.truncatingRemainder(dividingBy: 20.0) / 10
        print(pitch)
        return pitch < 0.5 ? pitch + 0.5 : pitch
    }
}
