import { persistStore, persistReducer } from "redux-persist";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { createStore, combineReducers } from "redux";

import reducer from "./reducer";

const rootReducer = combineReducers({
  user: reducer,
});

const persistConfig = {
  key: "user",
  storage: AsyncStorage,
};

const persistedReducer = persistReducer(persistConfig, rootReducer);

export let store = createStore(persistedReducer);
export let persistor = persistStore(store);
