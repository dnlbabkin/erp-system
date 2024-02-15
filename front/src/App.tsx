import React from 'react';
import './App.css';
import Header from './components/Header';
import { Container } from 'react-bootstrap';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import HomeScreen from './screens/HomeScreen';
import SignupScreen from './screens/SignupScreen';
import LoginScreen from './screens/LoginScreen';

function App() {
  return (
    <BrowserRouter>
        <Header />
        <main>
          <Container>
            <Routes>
              <Route path='/' element={<HomeScreen />} />
              <Route path='/signup' element={<SignupScreen />} />
              <Route path='/login' element={<LoginScreen />} />
            </Routes>
          </Container>
        </main>
      </BrowserRouter>
  );
}

export default App;
