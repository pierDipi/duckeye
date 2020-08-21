import React from 'react';
import './App.css';
import DuckTabs from "./components/tabs";
import {createMuiTheme, ThemeProvider} from "@material-ui/core/styles";

const theme = createMuiTheme({
    palette: {
        primary: {
            main: '#222'
        }
    }
});

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <ThemeProvider theme={theme}>
                    <DuckTabs/>
                </ThemeProvider>
            </header>
        </div>
    );
}

export default App;
