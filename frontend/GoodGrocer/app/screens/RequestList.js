import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, ScrollView, View, TouchableOpacity } from 'react-native';
import RequestCard from '../components/RequestCard';




const testData = [
  {
    name: "Angela",
    storeName: "Ralphs",
    numItems: 29, 
    imageUri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
    id: 1
  },
  {
    name: "Angela",
    storeName: "Ralphs",
    numItems: 30, 
    imageUri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
    id: 2
  },
  {
    name: "Angela",
    storeName: "Ralphs",
    numItems: 30, 
    imageUri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
    id: 3
  },
  {
    name: "Angela",
    storeName: "Ralphs",
    numItems: 30, 
    imageUri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
    id: 4
  },
  {
    name: "Angela",
    storeName: "Ralphs",
    numItems: 30, 
    imageUri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
    id: 5
  },
  {
    name: "Angela",
    storeName: "Ralphs",
    numItems: 30, 
    imageUri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
    id: 6
  }
]

function RequestList({setPage}) {
    const cards = testData.map((request) => <TouchableOpacity style={styles.requestCard} key={request.id} onPress={()=>setPage(1)}><RequestCard  name={request.name} storeName={request.storeName} numItems={request.numItems} imageUri={request.imageUri}/></TouchableOpacity>)
    return <SafeAreaView style={styles.container}>
    <View style={styles.content}>
      <Image source={require("../assets/logo.png")}/>
      <Text style={styles.titleText}>Requests in Westwood</Text>
      <ScrollView style={styles.listOfRequests}>
        {cards} 
      </ScrollView>
    </View>
  </SafeAreaView>; 

}


const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff'
    },
    content: {
      alignItems: 'center',
      marginTop: 40
    }, 
    listOfRequests: {
       marginBottom: 150,
       width: '85%%'
    },
    titleText:{
      fontSize: 25
    },
    requestCard: {
      paddingTop: 20
    }
  });

export default RequestList;