import React from 'react';
import { StackNavigator,TabNavigator, TabBarBottom } from 'react-navigation';
import VideoPage from '../components/VideoPage';
import Twitch from '../components/Video';
import Streams from '../components/Streams';
import Settings from '../components/Settings';
import Search from '../components/Search';
import { Ionicons } from '@expo/vector-icons';

export const StreamStack = StackNavigator(
  {
    Streams: { screen: Streams },
    VideoPage: {screen : VideoPage},
    Video : {screen : Twitch}
  },
  {
    navigationOptions: {
      headerStyle: { display: 'none' }
    }
  }
);



export const Tabs = TabNavigator(
  {
    Streams: { screen: StreamStack},
    Search: { screen: Search },
    Settings: { screen: Settings }
  },
  {
    navigationOptions: ({ navigation }) => ({
      title: navigation.state.routeName,
      tabBarIcon: ({ focused, tintColor }) => {
        const { routeName } = navigation.state;
        let iconName;
        if (routeName === 'Streams') {
          iconName = `logo-twitch`;
        } else if (routeName === 'Settings') {
          iconName = `ios-options${focused ? '' : '-outline'}`;
        } else if (routeName === 'Search') {
          iconName = `ios-search${focused ? '' : '-outline'}`;
        }
        return <Ionicons name={iconName} size={25} color={tintColor} />;
      }
    }),
    tabBarComponent: TabBarBottom,
    tabBarPosition: 'bottom',
    tabBarOptions: {
      activeTintColor: 'blue',
      inactiveTintColor: 'gray'
    },
    animationEnabled: true,
    swipeEnabled: true
  }
);



