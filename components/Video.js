import SystemJS from 'systemjs';
import {View} from 'react-native';

const TwitchPlayer = SystemJS.import('http://player.twitch.tv/js/embed/v1.js')

const player = () => {
    return (
        <View ref={ e => {handlePlayer()}}>
        </View>
    )
}