import React, {Component} from 'react';
import {View, WebView} from 'react-native';

export default class Twitch extends Component{
    render(){
        const login = this.props.navigation.state.params
        return (
                    <WebView
                        style={{ flex: 1 }}
                        javaScriptEnabled
                        mediaPlaybackRequiresUserAction={false}
                        source={{ uri: `https://player.twitch.tv/?channel=${login}` }}
                    />

        )
    }
} 