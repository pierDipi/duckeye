import React from 'react';

// Uncomment this line and the img reference
// import logo from './../logo.png';

import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Box from '@material-ui/core/Box';
import DuckGrid from "./list";

const host = window.location.host
const protocol = window.location.protocol
const path = "/clusterducktypes"

const url = `${protocol}//${host}${path}`

console.log(url)

function TabPanel(props) {
    const {children, value, index, ...other} = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`duck-tabpanel-${index}`}
            aria-labelledby={`duck-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box p={3}>
                    {children}
                </Box>
            )}
        </div>
    );
}

TabPanel.propTypes = {
    children: PropTypes.node,
    index: PropTypes.any.isRequired,
    value: PropTypes.any.isRequired,
};

function a11yProps(index) {
    return {
        id: `duck-tab-${index}`,
        'aria-controls': `duck-tabpanel-${index}`,
    };
}

export default class DuckTabs extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            ducks: [],
            value: 0,
        }

        const fetchClusterDuckTypes = () => {
            fetch(new Request(url, {method: "GET"}))
                .then(response => {
                    if (response.status >= 300) {
                        throw new Error(`Expected 2xx status code, got ${response.status}`)
                    }
                    return response.json()
                })
                .then(json => {

                    const sorted = json.sort((a, b) => a.name > b.name ? 1 : -1)

                    const newState = {
                        ...this.state,
                        ducks: sorted
                    }

                    this.setState(newState)

                })
                .catch(cause => console.debug(cause))
        }

        fetchClusterDuckTypes()

        // start polling cluster duck types
        setInterval(fetchClusterDuckTypes, 3000)
    }

    handleChange = (event, newValue) => {
        this.setState({
            ...this.state,
            value: newValue,
        })
    }

    render() {
        return (
            <div className="Tab">
                <AppBar position="absolute">
                    <Tabs
                        value={this.state.value}
                        onChange={this.handleChange}
                        aria-label="Duck tabs">

                        <Tab
                            label="Duckeye"
                            className="Tab-label"
                            {...a11yProps(0)}
                        />

                        {
                            this.state.ducks.map((duck, index) => {
                                return <Tab
                                    className="Tab-label"
                                    key={duck.name}
                                    label={duck.name}
                                    {...a11yProps(index + 1)}
                                />
                            })
                        }
                    </Tabs>
                </AppBar>

                <TabPanel
                    value={this.state.value}
                    index={0}
                >
                    <Box>
                        {/*<img src={logo} className="App-logo" alt="logo"/>*/}
                        <h1>Duckeye</h1>
                        <strong>
                            Discover Knative duck types
                        </strong>
                    </Box>
                </TabPanel>

                {
                    this.state.ducks.map((duck, index) => {
                        return <TabPanel
                            key={duck.name}
                            value={this.state.value}
                            index={index + 1}
                        >
                            <DuckGrid duck={duck}/>
                        </TabPanel>
                    })
                }
            </div>
        );
    }
}