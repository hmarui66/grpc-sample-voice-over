//
//  ContentView.swift
//  grpcvoiceover
//
//  Created by 丸井弘嗣 on 2020/02/28.
//  Copyright © 2020 hmarui66. All rights reserved.
//

import SwiftUI

struct ContentView: View {
    @ObservedObject(initialValue: ContentViewModel()) var vm: ContentViewModel
    @State private var token: String = ""

    var body: some View {
        VStack {
            if self.vm.isExecuting {
                text(self.vm.comment)
            } else {
                HStack {
                    TextField("input token", text: self.$token)
                        .textFieldStyle(RoundedBorderTextFieldStyle())
                        .padding()
                    Button(action: {
                        self.vm.subscribe(self.token)
                    }) {
                        Text("subscribe")
                    }.disabled(self.token == "")
                }
            }
        }
    }

    func text(_ c: Comment_Comment) -> Text {
        var font: Font = .system(size: 64)
        var color: Color = .blue
        switch self.vm.comment.text.count {
        case 0..<13:
            font = .system(size: 128)
            color = .red
        case 13..<52:
            font = .system(size: 64)
            color = .orange
        case 52..<100:
            font = .system(size: 32)
            color = .blue
        default:
            font = .system(size: 16)
            color = .black
        }
        return Text(self.vm.comment.text).font(font).foregroundColor(color)
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
