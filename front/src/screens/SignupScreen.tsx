import React, { SyntheticEvent, useState } from "react";
import { Form, Button, Modal } from "react-bootstrap";
import FormContainer from "../components/FormContainer";
import { useNavigate } from "react-router-dom";
import ModalWindow from "../components/ModalWindow";

const SignupScreen = () => {
    const [name, setName] = useState('')
    const [lastName, setLastName] = useState('')
    const [thirdName, setThirdName] = useState('')
    const [password, setPassword] = useState('')
    const [uid, setUid] = useState<any>()
    const [show, setShow] = useState(false)
    let navigate = useNavigate()

    const handleClose = () => {
        setShow(false)
        navigator.clipboard.writeText(uid)

        navigate("/login")
    }
    const handleShow = () => setShow(true) 


    const submitHandler = async(e: SyntheticEvent) => {
        e.preventDefault()
        
        const response = await fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name,
                last_name: lastName,
                third_name: thirdName,
                password
            })
        })

        const data = await response.json()
        setUid(data.uid)
        console.log(uid)

    }   
    
    return (
        <>
        <ModalWindow show={show} uid={uid} handleClose={handleClose} />
        <FormContainer>
            <Form onSubmit={submitHandler}>
                <Form.Group controlId="name" className="my-3">
                    <Form.Label>Имя</Form.Label>
                    <Form.Control type='name' placeholder="Введите ваше имя" 
                    value={name}
                    onChange={n => setName(n.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="lastName" className="my-3">
                    <Form.Label>Фамилия</Form.Label>
                    <Form.Control type='lastName' placeholder="Введите вашу фамилию" 
                    value={lastName}
                    onChange={l => setLastName(l.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="thirdName" className="my-3">
                    <Form.Label>Отчество</Form.Label>
                    <Form.Control type='thirdName' placeholder="Введите ваше отчество" 
                    value={thirdName}
                    onChange={t => setThirdName(t.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="password" className="my-3">
                    <Form.Label>Пароль</Form.Label>
                    <Form.Control type='password' placeholder="Введите ваш пароль" 
                    value={password}
                    onChange={p => setPassword(p.target.value)}
                    />
                </Form.Group>

                <Button variant="primary" type="submit" onClick={handleShow}>
                    Создать пользователя
                </Button>
            </Form>
        </FormContainer>
        </>
    )
}

export default SignupScreen