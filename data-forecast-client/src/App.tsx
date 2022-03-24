import React from 'react';
import logo from './logo.svg';
import './App.css';
import {useSelector} from 'react-redux';
import AuthModel from './models/authModel';
import { RootState } from './store/store';

import { useAppSelector, useAppDispatch } from './store/hooks'
import { Routes, Route, Outlet, Link } from "react-router-dom";

import Login from './pages/Login'
import Home from './pages/Home'

interface Props {
}

const App: React.FC<Props> = () => {

  // These work even better
  const authState = useAppSelector((state) => state.auth)
  const isAuth = useAppSelector((state) => state.auth.isAuthtenticated)
  console.log("check hook output: ", authState);
  const dispatch = useAppDispatch()

  // now is working!!! - but still missing something
  const isLoggedIn = useSelector<RootState>(state => state.auth.isAuthtenticated);

  console.log("-> isLoggedIn: ", isLoggedIn);

  return (
    <>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="home" element={<Home />} />
      </Routes>

      {/* <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="about" element={<About />} />
        <Route path="dashboard" element={<Dashboard />} />
        <Route path="*" element={<NoMatch />} />
      </Route>
      <Outlet /> */}
    </>
  );
}

export default App;
