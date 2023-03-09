import React, {useState} from 'react';
import { SafeAreaView, StyleSheet, Pressable, Image, Text, Title, View, ScrollView } from 'react-native';
import TextInput from '../components/TextInput';
import Button from '../components/Button';
import {Colors, Font} from '../Constants';
import useRequest from "../hooks/useRequest";

function ChangeName({navigation}) {
    const [name, setName] = useState('');

    const changeUserName = useRequest({
        url: "/user/update-name",
        method: "post",
        body: {
          name: name,
        },
        onSuccess: (data) => {
            console.log(name);
            props.navigation.navigate("YourCommunities");
        },
        onFail: (data) => {
            console.log(name);
        },
      });

    return (
        <SafeAreaView style={styles.container}>
            <View style={{marginTop: 20, marginLeft: 20}}>
                <Text style={styles.title}>Change your Name</Text>
            </View>
            <View style ={{ marginTop: 10, marginLeft: 30, marginRight: 30}}>
                <TextInput onChange={name => setName(name.nativeEvent.text)}></TextInput>
            </View>
            <View style={{alignItems: 'center'}}>
                <Button
                    title={"Submit"}
                    onPress={async () => await changeUserName.doRequest()}
                    textColor={"white"}
                    backgroundColor={Colors.lightGreen}
                    width={250}>
                </Button>
            </View>
        </SafeAreaView>

    );

}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  title: {
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
});

export default ChangeName;
