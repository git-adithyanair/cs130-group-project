import React, {useEffect, useState} from 'react';
import { SafeAreaView, StyleSheet, Text, Image, ScrollView, View } from 'react-native';
import ItemCard from '../components/ItemCard'
import useRequest from '../hooks/useRequest';

function RequestDetail(props) {
  const [items, setItems] = useState([]); 

  const getItemsInRequest = useRequest({
    url: `/request/items/${props.route.params.requestId}`,
    method: "get",
    onSuccess: (data) => {
      console.log(data)
     data.forEach((item)=>{setItems(oldArray=>[...oldArray, {itemName: item.name, numOfItem: item.quantity+item.quantity_type, id: item.id}])}); 
    }
  });
  const func = async () => getItemsInRequest.doRequest(); 
  useEffect(()=> {func()},[]); 


  const itemList = items.map((item) => <View style={styles.itemCard} key={item.id}><ItemCard itemName={item.itemName} numOfItem={item.numOfItem} /></View>);
    return <SafeAreaView style={styles.container}>
    <View style={styles.content}>
    <Image source={require("../assets/logo.png")}/>
    <Image style={styles.profileImage} source={{uri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png"}}/>
    <Text>Store: {props.route.params.storeName}</Text>
    <Text>Item Count: {items.length}</Text>
    <ScrollView style={styles.listOfItems}>
        {itemList}
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
    titleText:{
      fontSize: 25
    },
    subtitleText:{
      fontSize: 14
    },
    profileImage: {
      width: 75,
      height: 75,
      borderRadius: 75 / 2,
      marginTop: 15
    },
    itemCard: {
      marginTop: 20
    },
    listOfItems:{
      marginBottom: 270,
      width: '75%',
      marginTop: 15
    },
    titleCard: {
      display: 'flex',
      flexDirection: 'row',
      backgroundColor: 'red'
    }
  }); 

export default RequestDetail;