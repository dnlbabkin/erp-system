import React from "react";
import { Form, Button } from 'react-bootstrap';
import FormContainer from "../components/FormContainer";

const LoginScreen = () => {
    return (
        <FormContainer>
            <Form>
                <Form.Group controlId="name" className="my-3">
                    <Form.Label>Имя</Form.Label>
                    <Form.Control type='name' placeholder="Введите ваше имя" />
                </Form.Group>

                <Form.Group controlId="secondName" className="my-3">
                    <Form.Label>Фамилия</Form.Label>
                    <Form.Control type='secondName' placeholder="Введите вашу фамилию" />
                </Form.Group>

                <Form.Group controlId="thirdName" className="my-3">
                    <Form.Label>Отчество</Form.Label>
                    <Form.Control type='thirdName' placeholder="Введите ваше отчество" />
                </Form.Group>

                <Button variant="primary" type="submit">
                    Войти
                </Button>
            </Form>
        </FormContainer>
    )
}

export default LoginScreen