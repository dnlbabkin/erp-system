import React from "react";
import { Navbar, Nav, Container} from 'react-bootstrap'

const Header = () => {
    return (
        <Navbar bg='dark' variant='dark' expand="lg" collapseOnSelect >
            <Container>
                <Navbar.Brand href="/">erp</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="ms-auto">
                    <Nav.Link href="/signup">Регистрация</Nav.Link>
                    <Nav.Link href="/login">Вход</Nav.Link>
                </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default Header