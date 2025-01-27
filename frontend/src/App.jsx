import React from "react";
import Dashboard from "./components/Dashboard";
import SuspiciousActivity from "./components/SuspiciousActivity";

function App() {
  return (
    <div>
      <h1>Transaction Monitoring System</h1>
      <Dashboard />
      <SuspiciousActivity />
    </div>
  );
}

export default App;
