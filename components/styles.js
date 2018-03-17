import { StyleSheet } from 'react-native';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
    padding : 20
  },
  Heading: {
    fontSize: 36,
    marginTop: 50,
    marginLeft: 20,
    fontWeight: 'bold'
  },
  heroImage:{
    flex:1,
    width : "100%",
    height: 200,
    alignItems : "center",
    justifyContent:"center",
    padding: 5
  },
  featuredTitle:{
    fontSize: 30,
    fontWeight: 'bold',
    color:"#fff"
  },
  image: {
    height: 'auto',
    width: '85%'
  },
  roundedImage : {
    height: 600,
    width : '100%'
  },
  videoPreview : {
    height : 400,
    width : '100%',
    marginTop: 20
  }
});

export default styles;
