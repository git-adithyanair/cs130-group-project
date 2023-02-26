import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, ScrollView, View, TouchableOpacity } from 'react-native';
import RequestCard from '../components/RequestCard';





function Profile({setPage}) {
    return <SafeAreaView style={styles.container}>
    <View style={styles.content}>
      <Image source={require("../assets/logo.png")}/>
      <View style={styles.listOfRequests}>
        <View style={styles.leftColumn}>
            <Image style={styles.profileImage} source={{uri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png"}}/>      
        </View>
        <View style={styles.requestDetails}>
            <View><Text style={styles.titleText}>Angela</Text></View>
            <View><Text>Number of neighborhoods: 5</Text></View>
        </View>
      </View>
      <View style={styles.listOfChanges}>
          <View style={styles.changeItem}><Text style={styles.titleText}>Change Address</Text></View>
          <View style={styles.changeItem}><Text style={styles.titleText}>Change Name</Text></View>
          <View style={styles.changeItem}><Text style={styles.titleText}>Change Profile Picture</Text></View>
          <View style={styles.changeItem}><Text style={styles.titleText}>Join Community</Text></View>
          <View style={styles.changeItem}><Text style={styles.titleText}>Create Community</Text></View>
      </View>
    </View>
  </SafeAreaView>; 

}


const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  content: {
    alignItems: 'center',
    marginTop: 40
  },
  listOfRequests: {
    display: "flex",
    flexDirection: "row",
    paddingTop: 20
  },
  requestDetails:{
    flexDirection: "column",
    paddingLeft: 10,
    alignItems: 'center',
    marginTop: 10
  },
  profileImage: {
    width: 80,
    height: 75,
    borderRadius: 20
  },
  leftColumn:{
    textAlign: 'center'
  },
  titleText: {
    fontSize: 25
  },
  listOfChanges: {
    marginTop: 80,
    height: '60%',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
  }, 
  changeItem: {
    borderTopWidth: 2,
    width: 300
  }
});

export default Profile;