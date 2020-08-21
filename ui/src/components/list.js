import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import {Grid} from "@material-ui/core";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";

const useStyles = makeStyles(() => ({
    root: {
        width: '100%',
        height: "90%",
        marginTop: "10%"
    },
}));

export default function DuckGrid(props) {

    const {duck} = props

    const classes = useStyles();

    if (duck.status === undefined || duck.status == null) {
       return (
           <div>
               <h3>No items found</h3>
           </div>
       )
    }

    const items = Object.keys(duck.status).flatMap((key) => {
        return duck.status[key].flatMap(value => {
            return {
                version: key,
                ...value
            }
        })
    })

    items.sort((a, b) => a.kind > b.kind ? 1 : -1)

    return (
        <div className={classes.root}>
            <Grid container spacing={3}>
                {
                    items.map(item => {
                        return <Grid item key={item.apiVersion + item.kind}>
                            <Card>
                                <CardContent>
                                    <Typography className={classes.title} color="textSecondary" gutterBottom>
                                        {item.apiVersion}
                                    </Typography>
                                    <Typography variant="h5" component="h2">
                                        {item.kind}
                                    </Typography>
                                    <Typography className={classes.pos} color="textSecondary">
                                        Scope: {item.scope}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                    })
                }
            </Grid>
        </div>
    );
}