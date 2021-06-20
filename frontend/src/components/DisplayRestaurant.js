import React from "react";
import {Grid, Typography, makeStyles} from "@material-ui/core";
const useStyles = makeStyles((theme)=>({
    paper: {
      margin: "10px 5px",
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      
      borderRadius:"5px",
      backgroundColor:'white'
    },
    
}));

function MyCard(props){
    const classes = useStyles();
    const item = props.info;
    return <div className={classes.root}>
    <Grid item   className={classes.paper}  >
        <Grid container spacing={1} >
            <Grid item xs={4} style={{padding:"10px"}}>
                <img src="https://www.chelseasmessyapron.com/wp-content/uploads/2014/02/Red-and-White-Pasta-3.jpg" width="100%" height="100%"></img>
            </Grid>
            <Grid item xs={8}>
                
                        <Grid container>
                            <Grid item xs={10}>
                                <Typography component="h6" variant="h4"><strong>{item.name}</strong></Typography>
                                <div style={{color:"grey"}}>
                                    Serves: {item.cuisine}
                                </div>
                                <div style={{color:"grey"}}>
                                    Located at: {item.address}
                                </div>
                                <div style={{color:"grey"}}>
                                    CFT: ₹{item.costForTwo}
                                </div>
                                <div style={{marginTop: "20px"}}> 
                                    <h6>Open from {item.openingTime} to {item.closingTime}</h6>
                                </div>
                            

                            </Grid>
                            
                                
                        </Grid>
                        <div style={{display: "flex", flexDirection: "column", marginTop: "90px"}}>
                    <div style={{color:"red"}}>
                       <h5> Rating: {item.rating}</h5>
                    </div>
                    <div>
                    </div>

                </div>
            </Grid>
        </Grid>
    </Grid>
    </div>
}

export default MyCard;