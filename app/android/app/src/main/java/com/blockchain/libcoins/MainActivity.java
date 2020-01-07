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

        // Bitcoin signature
        String btcecKey = "22a47fa09a223f2aa079edf85a7c2d4f8720ee63e502ee2869afab7de234b80c";
        BitcoinSignResult btcResult = Libcoins.bitcoinSign(btcecKey, message);

        // Ethereum signature
        String ecdsaKey = "127e668421cbbf6a80692679560c618d5f06281b02a8323157816e4c7ce50e2b3b776e2f6d9febc03d3abdf91c16d15396dc6f72bec3e259df2bfdec8fe41f89";
        String ecdsaSignature = Libcoins.ethereumSign(ecdsaKey, message);

        // Ethereum signature
        String text = String.format("Bitcoin signature: %s\n\nEthereum signature:%s", btcResult.unwrap(), ecdsaSignature);
        mTextView.setText(text);
    }
}
