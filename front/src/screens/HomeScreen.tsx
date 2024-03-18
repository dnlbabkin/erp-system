import React from "react";
import LoginScreen from "./LoginScreen";
import SignupScreen from "./SignupScreen";
import TransactionScreen from "./TransactionScreen";

interface Props {
    uid: any
}

const HomeScreen = ({ uid }: Props) => {
    return uid ? (
        <TransactionScreen />
    ) : (
        <SignupScreen />
    )
}

export default HomeScreen