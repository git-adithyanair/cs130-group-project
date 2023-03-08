import React, {useState} from 'react';
import {SafeAreaView, StyleSheet, FlatList, Pressable, Image, Text, Title, View, ScrollView } from 'react-native';
import Login from './Login';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import Modal from "react-native-modal";
import Button from '../components/Button';
import TextInput from '../components/TextInput';
import ItemCard from '../components/ItemCard';
import {Colors, Font} from '../Constants';
import { Picker } from '@react-native-picker/picker';



const Tab = createBottomTabNavigator();



function Buy({navigation, route}) {
    const [store, setStore] = useState('');
    const [item, setItem] = useState('');
    const [numItems, setNumItems] = useState('');
    const [type, setType] = useState('');
    const [brand, setBrand] = useState('');
    const [notes, setNotes] = useState('');
    const [items, setItems] = useState([]);
    const [isModalVisible, setIsModalVisible] = React.useState(false);
    const handleModal = () => setIsModalVisible(() => !isModalVisible);
    const [DATA, setDATA] = useState([]);

    const Item = ({title}) => (
      <View style={styles.item}>
        <Text style={styles.title}>{title}</Text>
      </View>
    );

    const addItem = () => {
      setIsModalVisible(() => !isModalVisible);
      let dataPoint = {id: '5s694d0f-3da1-4r1f-bd96-145571e29d72',
      title: 'Fourth Item'};
      setDATA(prev => [...prev, dataPoint]);

      console.log(DATA);
      let individualItem = {"name": item,
                            "quantity_type": type,
                            "quantity": numItems,
                            "preferred_brand": brand,
                            "image": null,
                            "extra_notes": notes,};

      setItems(prev => [...prev, individualItem]);
    }
    const completeOrder = () => {
        navigation.navigate('OrderCreated');
      }
    return (
        <SafeAreaView style={styles.container}>
          <View style={{marginTop: 10, marginBottom: 30}}>
                <View style={{marginLeft: 20}}>
                    <Text style={styles.title}>Create an Order in {route.params.communityName}</Text>
                    <Text style={{fontSize: 18, marginTop: 5}}>Pick your Store </Text>
                </View>
                <View style={{marginTop: 10, marginLeft: 30, marginRight: 30}}>
                    <TextInput onChange={store => setStore(store.nativeEvent.text)}/>
                </View>
                <View style={styles.content}>
                    <Button title={"Add Items"} onPress={handleModal} textColor={"white"} backgroundColor={Colors.lightGreen} width={250} />
                    <Modal isVisible={isModalVisible} transparent={true} style={styles.modalStyle}>
                    <View
                            style={styles.outerModal}>
                            <View
                                style={styles.innerModal}>
                                <View style={{marginTop: 10}}>
                                    <Text style={styles.modalFont}>Type of Item</Text>
                                </View>
                                <View style={styles.modalTextinput}>
                                  <TextInput onChange={item => setItem(item.nativeEvent.text)}/>
                                </View>
                                <View>
                                    <View style={styles.content}>
                                      <Text style={styles.modalFont}>Quantity Type</Text>
                                    </View>
                                    <Picker
                                      selectedValue={type}
                                      onValueChange={(itemValue, itemIndex) =>
                                        setType(itemValue)
                                      }>
                                      <Picker.Item label="lbs" value="lbs" />
                                      <Picker.Item label="oz" value="oz" />
                                      <Picker.Item label="fl_oz" value="fl_oz" />
                                      <Picker.Item label="gal" value="gal" />
                                      <Picker.Item label="litres" value="litres" />
                                    </Picker>
                                </View>
                                <View>
                                    <Text style={styles.modalFont}>
                                        Quantity of Item
                                    </Text>
                                    <View style={styles.modalTextinput}>
                                      <TextInput onChange={numItems => setNumItems(numItems.nativeEvent.text)}/>
                                    </View>
                                </View>
                                <View>
                                    <Text style={styles.modalFont}>
                                        Preferred Brand
                                    </Text>
                                    <View style={styles.modalTextinput}>
                                      <TextInput onChange={brand => setBrand(brand.nativeEvent.text)}/>
                                    </View>
                                </View>
                                <View>
                                    <Text style={styles.modalFont}>
                                        Notes
                                    </Text>
                                    <View style={styles.modalTextinput}>
                                      <TextInput onChange={notes => setNotes(notes.nativeEvent.text)}/>
                                    </View>
                                </View>
                                <View style={{ alignItems: 'center', marginBottom: 10}}>
                                    <Button
                                      title={"Add Item"}
                                      onPress={addItem}
                                      textColor={"white"}
                                      backgroundColor={Colors.lightGreen}
                                      width={200}/>
                                </View>

                        </View>
                    </View>
                    </Modal>
                </View>
            </View>
          <View style={{marginLeft: 20, marginBottom: 10}}>
            <Text style={styles.title}>Your Items</Text>
          </View>
          <View>
          <FlatList
            data={DATA}
            renderItem={({item}) => <Item title={item.title} />}
            // extraData={changedData}
            keyExtractor={item => item.id}
          />
          </View>
            {/* <ScrollView style={{margin: 20}}> */}
                {/* <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard>
                <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard>
                <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard>
                <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard> */}
                <View style={{alignItems: 'center', marginTop: 20, marginBottom: 20}}>
                    <Button
                      title={"Complete your Order"}
                      onPress={() => completeOrder()}
                      textColor={"white"}
                      backgroundColor={Colors.blue}
                      width={300} />
                </View>
            {/* </ScrollView> */}
        </SafeAreaView>
    );

}


const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff',
    },
    content: {
      alignItems: 'center'
    },
    title: {
      fontSize: Font.s1.size,
      fontFamily: Font.s1.family,
      fontWeight: Font.s1.weight,
    },
    modalFont: {
      fontSize: Font.s3.size,
      fontFamily: Font.s3.family,
      fontWeight: Font.s3.weight,
      marginLeft: 20,
    },
    modalStyle: {
      backgroundColor: Colors.cream,
      margin: 50,
    },
    outerModal: {
      backgroundColor: Colors.lightGreen,
      alignItems: 'center',
      justifyContent: 'center',
    },
    innerModal: {
      // alignItems: 'center',
      backgroundColor: Colors.cream,
      marginVertical: 60,
      width: '90%',
      borderWidth: 1,
      borderColor: '#fff',
      borderRadius: 7,
      elevation: 6,
    },
    modalTextinput: {
      marginLeft: 20,
      marginRight: 20,
    },
    item: {
      backgroundColor: '#f9c2ff',
      padding: 20,
      marginVertical: 8,
      marginHorizontal: 16,
    },
  });

export default Buy;