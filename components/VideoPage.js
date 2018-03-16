import React, { Component } from "react";
import {View, Text, WebView, Image} from 'react-native'
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import styles from "./styles";
import {fetchVideoURL} from '../actions';
// import Video from 'react-native-video'

class VideoPage extends Component{

    constructor(props){
        super(props);
        this.state = {
            User : {}
        }
    }

    componentDidMount(){
        const {fetchVideoURL} = this.props;
        fetchVideoURL(this.props.navigation.state.params)
    }

    componentWillReceiveProps(props){
        if(props.User){
            this.setState({User : props.User.user})
        }
    }

    componentWillUnmount(){
        this.state = {}
    }

    renderPage(){
        if(this.state.User.login){
            return (
                <View style={styles.container}>
                    <Image style={styles.roundedImage} resizeMethod="auto" resizeMode="cover" source={{uri : this.state.User.offline_image_url || "http://dotawallpaper.org/wp-content/uploads/2016/05/Necrophos%20Hero%20Dota%202.jpg"}} />
                    <Text style={styles.Heading}>
                            { (this.state.User.display_name).toUpperCase()}
                    </Text>
                    <WebView
                        style={styles.videoPreview}
                        source={{uri:`http://player.twitch.tv?channel=${this.state.User.login}`}}
                        mediaPlaybackRequiresUserAction
                        startInLoadingState
                        scalesPageToFit
                    />
                </View>
            )
        }
    }

    render(){
        return(
            <View>
                {this.renderPage()}
            </View>
        )
    }
}


 function mapStateToProps({User}){
    return {User}
}
function mapStateToDispatch(dispatch){
    return bindActionCreators({fetchVideoURL},dispatch)
}

export default connect(mapStateToProps,mapStateToDispatch)(VideoPage)