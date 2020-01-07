// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#import "ViewController.h"
@import Libcoins;  // Gomobile bind generated framework

@interface ViewController ()
@end

@implementation ViewController

@synthesize textLabel;

- (void)loadView {
    [super loadView];
    LibcoinsBitcoinSignResult *btcSig = LibcoinsBitcoinSign(@"22a47fa09a223f2aa079edf85a7c2d4f8720ee63e502ee2869afab7de234b80c", @"Hello, world!");
    NSString *text1 = [NSString stringWithFormat:@"Bitcoin signature: %@", [btcSig unwrap]];
    
    NSString *ethSig = LibcoinsEthereumSign(@"fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19", @"Hello, world!");
    NSString *text2 = [NSString stringWithFormat:@"Ethereum signature: %@", ethSig];
    textLabel.text = [NSString stringWithFormat:@"%@\n\n%@", text1, text2];
}

@end
