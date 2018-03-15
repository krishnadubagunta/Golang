import React, { Component } from 'react';
import { StyleSheet, Text, View } from 'react-native';
import styles from './styles';

export default class Streams extends Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {}

  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.Heading}>Streams</Text>
      </View>
    );
  }
}
