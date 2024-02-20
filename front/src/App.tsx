import React, { useEffect, useState } from 'react';
import './App.css';
import Header from './components/Header';
import { Container, Modal } from 'react-bootstrap';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import HomeScreen from './screens/HomeScreen';
import SignupScreen from './screens/SignupScreen';
import LoginScreen from './screens/LoginScreen';

function App() {
  const [uid, setUid] = useState<any>()

  useEffect(() => {
    ;(async () => {
      const response = await fetch('http://localhost:8080/api/user', {
        headers: { 'Content-Type': 'application/json' },
        credentials: "include"
      })

      const data = await response.json()
      setUid(data.uid)
    })()
  })
  
  return (
    <BrowserRouter>
        <Header uid={uid} setUid={setUid}/>
        <main>
          <Container>
            <Routes>
              <Route path='/' element={<HomeScreen uid={uid} />} />
              <Route path='/signup' element={<SignupScreen />} />
              <Route path='/login' element={<LoginScreen />} />
            </Routes>
          </Container>
        </main>
      </BrowserRouter>
  );
}

export default App;
