import React, {Component} from 'react';
import {View, WebView} from 'react-native';

export default class Twitch extends Component{
    render(){
        const login = this.props.navigation.state.params
        return (
                <View style={{ flex: 1 }}>
                    <WebView
                        style={{ flex: 2 }}
                        javaScriptEnabled
                        mediaPlaybackRequiresUserAction={false}
                        source={{ uri: `https://player.twitch.tv/?channel=${login}` }}
                    />
                </View>
        )
    }
} 