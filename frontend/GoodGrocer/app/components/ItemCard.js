import React from "react";
import {Image, Text ,View, StyleSheet } from 'react-native';
import {Card, Button , Title ,Paragraph } from 'react-native-paper';
import { Dim, Colors, Font, BorderRadius } from "../Constants";

const RequestCard = (props) => {
    return(

        <Card style={Styles.container}>
            <Card.Content>
                <Text>
                    <Title style={{marginTop: 20}}>Item: </Title>
                    <Text>{props.itemName}</Text>
                </Text>
                <Text>
                    <Title style={{textAlign: 'right'}}>Amount: </Title>
                    <Text style={{textAlign: 'right'}}>{props.numOfItem}</Text>
                </Text>
            </Card.Content>
        </Card>


    );

};


const Styles = StyleSheet.create({
    container :{
        backgroundColor: Colors.white,
        alignContent:'center',

    },
    text: {
        marginTop:30,
    }
});


export default RequestCard;

