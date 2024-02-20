import React, { SyntheticEvent } from "react";
import { Navbar, Nav, Container} from 'react-bootstrap'

interface Props {
    uid: any
    setUid: (uid: any) => void
}

const Header = ({ uid, setUid }: Props) => {

    const logoutHandler = async(e: SyntheticEvent) => {
        e.preventDefault()

        await fetch('http://localhost:8080/api/logout', {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
            credentials: "include", 
        })

        setUid(0)
    }

    return (
        <Navbar bg='dark' variant='dark' expand="lg" collapseOnSelect >
            <Container>
                <Navbar.Brand href="/">erp</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">

                    {uid ? (
                        <Nav className="ms-auto">
                            <Nav.Link>{uid}</Nav.Link>
                            <Nav.Link onClick={logoutHandler}>Выход</Nav.Link>
                        </Nav>
                    ) : (
                        <Nav className="ms-auto">
                            <Nav.Link href="/signup">Регистрация</Nav.Link>
                            <Nav.Link href="/login">Вход</Nav.Link>
                        </Nav>
                    )}
                
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default Header