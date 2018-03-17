import React, { Component } from "react";
import {View, Text, WebView, Image, ScrollView} from 'react-native'
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import styles from "./styles";
import {fetchVideoURL} from '../actions';
import { Button } from "react-native-elements";
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
            const {User} = this.state;
            return (
                <View style={styles.container}>
                    <Image style={styles.roundedImage} resizeMethod="auto" resizeMode="cover" source={{uri : User.offline_image_url || "http://dotawallpaper.org/wp-content/uploads/2016/05/Necrophos%20Hero%20Dota%202.jpg"}} />
                    <Text style={styles.Heading}>
                            { (User.display_name).toUpperCase()}
                    </Text>
                    <Button icon={{ name: 'video-label' }}
                            backgroundColor="#03A9F4"
                            buttonStyle={{
                                borderRadius: 0,
                                marginLeft: 0,
                                marginRight: 0,
                                marginBottom: 0,
                            }}
                            onPress={this.handleClick.bind(this)}
                            title="VIEW NOW"/>
                </View>
            )
        }
    }

    handleClick(){
        this.props.navigation.navigate('Video',this.state.User.login)
    }

    render(){
        return(
            <ScrollView>
            <View>
                {this.renderPage()}
            </View>
            </ScrollView>
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