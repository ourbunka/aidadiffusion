import { ChakraProvider, Button, Box, Center, Input, Text, 
  Grid, GridItem, Image, CircularProgress, Tooltip, Link } from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import { ColorModeSwitcher } from './Components/ColorswitchingComponent';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from './Components/Home';

function App() {
  return (
    
  <ChakraProvider>
    <Router basename={process.env.PUBLIC_URL}>
    <Routes>
      <Route path="/" element={<Home/>}></Route>
    </Routes>
    </Router>
    </ChakraProvider>
  );
}

export default App;
