package com.blockchain.libcoins;

import android.app.Activity;
import android.os.Bundle;
import android.widget.TextView;

import libcoins.Libcoins;
import libcoins.BitcoinSignResult;

public class MainActivity extends Activity {

    private TextView mTextView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        mTextView = (TextView) findViewById(R.id.mytextview);

        String message = "Hello world";
        String btcecKey = "22a47fa09a223f2aa079edf85a7c2d4f8720ee63e502ee2869afab7de234b80c";
        String ecdsaKey = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19";

        // Ethereum signature
        String text = String.format(
            "Bitcoin signature: %s\n\nEthereum signature: %s\n\nError handling: %s"
            , Libcoins.bitcoinSign(btcecKey, message).unwrap()
            , Libcoins.ethereumSign(ecdsaKey, message)
            , Libcoins.bitcoinSign("invalid key", message).unwrap()
        );
        mTextView.setText(text);
    }
}
