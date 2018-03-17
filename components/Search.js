import React, { Component } from 'react';
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native';
import styles from './styles';
import {SearchBar, List, ListItem} from 'react-native-elements';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import {fetchUsers} from '../actions';

class Search extends Component {
  constructor(props){
    super(props)
    this.state = {
      users:[],
      value: ''
    }
    this.handleUserSelect = this.handleUserSelect.bind(this);
  }

  handleChangeText(data){
    this.props.fetchUsers(data)
  }

  componentWillReceiveProps({users}){
    this.setState({
      users
    })
  }

  handleUserSelect(e){
    console.log(e);
  }

  render() {
    const {users} = this.state;
    return (
      <View>
        <Text style={styles.Heading}>Search</Text>
        <SearchBar lightTheme onChangeText={this.handleChangeText.bind(this)} icon={{type: 'material', color: '#86939e', name: 'search'}} round showLoadingIcon={ this.state.value != "" ? true : false } />
        <List containerStyle={{marginBottom: 20}}>
          {
            users.map((l, i) => (
              <TouchableOpacity key={i} onPress={e => this.handleUserSelect(e)}>
                <ListItem
                  roundAvatar
                  avatar={{uri:l.ProfileImage}}
                  title={l.DisplayName}
                />
              </TouchableOpacity>
            ))
          }
        </List>
      </View>
    );
  }
}

mapStateToProps = ({searchResults}) => {
  return {users : searchResults}
}

mapDispatchToProps = dispatch => {
  return bindActionCreators({fetchUsers},dispatch)
}

export default connect(mapStateToProps,mapDispatchToProps)(Search)