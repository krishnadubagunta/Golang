import React, { Component } from 'react';
import { View, Text, ImageBackground, TouchableOpacity } from 'react-native';
import { Card, ListItem, Button } from 'react-native-elements';
import { Ionicons } from '@expo/vector-icons';
import { replace } from 'lodash';
import styles from './styles';
import { connect } from 'react-redux';
import {selectID} from '../actions';
import {bindActionCreators} from 'redux';


class StreamCard extends Component {

  handleClick() {
    this.props.selectID(this.props.stream.user_id)
  }

  render() {
    const { stream } = this.props;
    var image = stream.thumbnail_url;
    image = replace(image, '{width}x{height}', '500x300');
    return (
      <Card style={{paddingBottom: 5}} featuredTitle={stream.title} image={{ uri: image }}>
  
        <Button
            icon={{ name: 'video-label' }}
            backgroundColor="#03A9F4"
            buttonStyle={{
              borderRadius: 0,
              marginLeft: 0,
              marginRight: 0,
              marginBottom: 0,
            }}
            onPress={this.handleClick.bind(this)}
            title="VIEW NOW"
            />
      </Card>)
  }
}

function mapStateToDispatch(dispatch){
  return bindActionCreators({selectID},dispatch);
}

export default connect(null,mapStateToDispatch)(StreamCard)