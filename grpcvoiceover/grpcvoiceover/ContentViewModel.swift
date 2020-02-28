//
//  ContentViewModel.swift
//  grpcvoiceover
//
//  Created by 丸井弘嗣 on 2020/02/28.
//  Copyright © 2020 hmarui66. All rights reserved.
//

import GRPC
import NIO
import Combine

final class ContentViewModel: ObservableObject {
    func subscribe() {
        // See: https://github.com/apple/swift-nio#eventloops-and-eventloopgroups
        let group = MultiThreadedEventLoopGroup(numberOfThreads: 1)

        // Make sure the group is shutdown when we're done with it.
        defer {
            try! group.syncShutdownGracefully()
        }

        // Provide some basic configuration for the connection, in this case we connect to an endpoint on
        // localhost at the given port.
        let configuration = ClientConnection.Configuration(
            target: .hostAndPort("10.36.27.215", 8080),
            eventLoopGroup: group
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
        }

        let status = try! call.status.recover { _ in .processingError }.wait()
        if status.code != .ok {
            print("PRC failed: \(status)")
        }

    }
}
