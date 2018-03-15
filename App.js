import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { StackNavigator } from 'react-navigation';
import HomeScreen from './Navigation/TabNavigation';

const AppScreens = StackNavigator(
  {
    Home: { screen: HomeScreen }
  },
  {
    navigationOptions: {
      headerStyle: { display: 'none' },
      headerLeft: null
    }
  }
);

export default class App extends React.Component {
  render() {
    return <AppScreens />;
  }
}
