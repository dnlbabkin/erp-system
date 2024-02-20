import React from "react";
import LoginScreen from "./LoginScreen";
import SignupScreen from "./SignupScreen";

interface Props {
    uid: any
}

const HomeScreen = ({ uid }: Props) => {
    return uid ? (
        <h1>{uid}</h1>
    ) : (
        <SignupScreen />
    )
}

export default HomeScreen