import React from "react";
import { render } from "react-dom";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./index.css";
import { TopStoriesPage } from "./TopStoriesPage";

render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<TopStoriesPage />}></Route>
    </Routes>
  </BrowserRouter>,
  document.getElementById("root")
);
