import React, { Component } from 'react';
import {  Text, View, FlatList, ScrollView } from 'react-native';
import { Card, ListItem, Button } from 'react-native-elements';
import styles from './styles';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { fetchStreams } from '../actions';
import StreamCard from './streamCard';
import { replace } from 'lodash';


class Streams extends Component {
  constructor(props) {
    super(props);
    this.renderStreams = this.renderStreams.bind();
  }

  componentDidMount() {
    this.props.fetchStreams();
  }

  renderStreams(streams) {
    if (streams) {
      return (
        <ScrollView>
          <Text style={styles.Heading}>Streams</Text>
          <FlatList
            data={streams}
            keyExtractor={(item, index) => item.id}
            renderItem={({ item }) => <StreamCard key={item.id} stream={item}/>} 
          />
        </ScrollView>
      );
    }
  }

  componentWillReceiveProps({ID}){
    if( ID ){
        this.props.navigation.navigate('Video',ID.ID)
    }
  }

  componentWillUnmount(){
    this.props.streams = []
  }

  render() {
    const { streams } = this.props;
    return <View>{this.renderStreams(streams)}</View>;
  }
}

function mapStateToProps({ streams, pagination, error, ID }) {
  return { streams: streams.streams, pagination, error, ID };
}
function mapStateToDispatch(dispatch) {
  return bindActionCreators({ fetchStreams }, dispatch);
}

export default connect(mapStateToProps, mapStateToDispatch)(Streams);
