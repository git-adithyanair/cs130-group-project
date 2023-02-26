import React, {useState} from 'react';
import { SafeAreaView, StyleSheet, Text, Image, ScrollView, View, TouchableOpacity } from 'react-native';

import Errand from './Errand'
import RequestList from './RequestList';
function Shop({navigation}) {
    const [page, setPage] = useState(0); 

    return page == 0 ? <RequestList setPage={setPage}/> : <Errand setPage={setPage}/>; 

}



export default Shop;