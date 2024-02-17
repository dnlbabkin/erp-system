import React from "react";

interface Props {
    name: String
}

const HomeScreen = ({name}: Props) => {
    return name ? (
        <h1>name</h1>
    ) : (
        <h1>Тут чета будет</h1>
    )
}

export default HomeScreen