import React, {useState} from 'react';
import { SafeAreaView, StyleSheet, Pressable, Image, Text, Title, View, ScrollView } from 'react-native';
import Modal from "react-native-modal";
import Button from '../components/Button';
import TextInput from '../components/TextInput';
import ItemCard from '../components/ItemCard';
import {Colors, Font} from '../Constants';

function CreateOrder({navigation}) {
    const [item, setItem] = useState('');
    const [numItems, setNumItems] = useState('');
    const [store, setStore] = useState('');
    const [isModalVisible, setIsModalVisible] = React.useState(false);
    const handleModal = () => setIsModalVisible(() => !isModalVisible);

    const completeOrder = () => {
        navigation.navigate('OrderCreated');
      }
    return (
        <SafeAreaView style={styles.container}>
            <View style={{alignItems: 'center'}}>
                <Image source={require("../assets/logo.png")}/>
            </View>
            <View style={{marginTop: 10, marginBottom: 30}}>
                <View style={{marginLeft: 20}}>
                    <Text style={styles.title}>Create an Order in </Text>
                    <Text style={{fontSize: 18, marginTop: 5}}>Pick your Store </Text>
                </View>
                <View style={{marginTop: 10, marginLeft: 30, marginRight: 30}}>
                    <TextInput onChange={store => setStore(store.nativeEvent.text)}/>
                </View>
                <View style={{alignItems: 'center'}}>
                    <Button title={"Add Items"} onPress={handleModal} textColor={"white"} backgroundColor={Colors.lightGreen} width={250} />
                    <Modal isVisible={isModalVisible} transparent={true} style={styles.modalStyle}>
                    <View
                            style={{
                                flex: 1,
                                backgroundColor: 'rgba(52, 52, 52, 0)',
                                alignItems: 'center',
                                justifyContent: 'center',
                            }}>
                            <View
                                style={{
                                alignItems: 'center',
                                backgroundColor: 'white',
                                marginVertical: 60,
                                width: '90%',
                                borderWidth: 1,
                                borderColor: '#fff',
                                borderRadius: 7,
                                elevation: 10,
                                }}>
                                <View style={{alignItems: 'center', margin: 10}}>
                                    <Text style={{fontSize: 25, marginTop: 5,}}>Type of Item</Text>
                                </View>
                                <View>
                                    <TextInput onChange={item => setItem(item.nativeEvent.text)}/>
                                    <Text style={{fontSize: 18, marginTop: 5, color:'grey'}}>
                                        Quantity of Item
                                    </Text>
                                    <TextInput onChange={numItems => setNumItems(numItems.nativeEvent.text)}/>
                                </View>
                                <View style={{margin: 10}}>
                                    <Button title={"Add Items"} onPress={handleModal} textColor={"white"} backgroundColor={Colors.lightGreen} width={200}/>
                                </View>

                        </View>
                    </View>
                    </Modal>
                </View>
            </View>
            <View style={{marginLeft: 20}}>
                <Text style={styles.title}>Your Items</Text>
            </View>
            <ScrollView style={{margin: 20}}>
                <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard>
                <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard>
                <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard>
                <ItemCard itemName={"Apples"} numOfItem={"3"}></ItemCard>
                <View style={{alignItems: 'center'}}>
                    <Button title={"Complete you Order"} onPress={() => completeOrder()} textColor={"white"} backgroundColor={Colors.blue} width={300} />
                </View>
            </ScrollView>
        </SafeAreaView>

    );

}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    marginBottom: 50,
    // alignItems: 'center',
  },
  title: {
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
  modalStyle: {
    backgroundColor: 'F3F3F3',
    // padding: 20,
    margin: 50
  },
});

export default CreateOrder;