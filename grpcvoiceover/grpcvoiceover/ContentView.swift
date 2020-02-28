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
    var body: some View {
        HStack {
            Text("Hello, World!")
            Button(action: {
                self.vm.subscribe()
            }) {
                Text("subscribe")
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
