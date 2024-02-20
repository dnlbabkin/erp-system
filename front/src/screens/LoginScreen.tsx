import React, { SyntheticEvent, useState } from "react";
import { Form, Button } from 'react-bootstrap';
import FormContainer from "../components/FormContainer";
import { useNavigate } from "react-router-dom";

const LoginScreen = () => {
    const [uid, setUid] = useState<any>()
    const [password, setPassword] = useState('')
    let navigate = useNavigate();

    const submitHandler = async(e: SyntheticEvent) => {
        e.preventDefault()
        
        await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: "include", 
            body: JSON.stringify({
                uid,
                password
            })
        })

        console.log(uid)

        navigate('/')
    }

    return (
        <>
        <FormContainer>
            <Form onSubmit={submitHandler}>
                <Form.Group controlId="uid" className="my-3">
                    <Form.Label>Идентификационный id</Form.Label>
                    <Form.Control type='uid' placeholder="Введите ваш идентификационный id"
                        value={uid}
                        onChange={u => setUid(parseInt(u.target.value))} />
                </Form.Group>

                <Form.Group controlId="password" className="my-3">
                    <Form.Label>Пароль</Form.Label>
                    <Form.Control type='password' placeholder="Введите ваш пароль"
                        value={password}
                        onChange={p => setPassword(p.target.value)} />
                </Form.Group>

                <Button variant="primary" type="submit">
                    Войти
                </Button>
            </Form>
        </FormContainer></>
    )
}

export default LoginScreen