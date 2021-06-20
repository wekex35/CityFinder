import fetch from 'cross-fetch';
import React from 'react';
import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete';
import CircularProgress from '@material-ui/core/CircularProgress';
import './mapbox-gl.css';
import ReactMapboxGl, { Layer, Feature } from 'react-mapbox-gl';

function sleep(delay = 0) {
    return new Promise((resolve) => {
        setTimeout(resolve, delay);
    });
}

export default function AutoSearch() {
    const [open, setOpen] = React.useState(false);
    const [options, setOptions] = React.useState([]);
    const [cities, setCities] = React.useState([]);
    const loading = open;
    const Map = ReactMapboxGl({
        accessToken:
            'pk.eyJ1Ijoid2VrZXgzNSIsImEiOiJja3E0bmF3emgwM2FkMnBtdHM5a3FkaWo5In0.8pjWQgGpx9guRNgC2LfJ6g'
    });

    React.useEffect(() => {
        if (!open) {
            setOptions([]);
        }

    }, [open]);



    const onChange = (event, value) => {
        let active = true;
        fetch(`http://localhost:7007/adjacent-city?city=${value}`).then(response => {
            const res = response.json().then((a) => {
                console.log(a.data);
                setCities(a.data)

                // if (active) {
                //     setOptions(a.data);
                // }
            });
        });
        return () => {
            active = false;
        };
    };

    const handleChange = (event, value) => {
        console.log(value);
        console.log(loading);


        let active = true;

        fetch(`http://localhost:7007/search-city?query=${value}`).then(response => {
            const res = response.json().then((a) => {
                console.log(a);

                if (active) {
                    setOptions(a.data);
                }
            });


        });

        return () => {
            active = false;
        };

    };

    return (
        <>
            <Autocomplete
                id="asynchronous-demo"
                style={{ width: 300 }}
                open={open}
                onOpen={() => {
                    setOpen(true);
                }}
                onClose={() => {
                    setOpen(false);
                }}
                getOptionSelected={(option, value) => option === value}
                getOptionLabel={(option) => option}
                options={options}
                loading={loading}
                onChange={onChange}
                onInputChange={handleChange}
                renderInput={(params) => (
                    <TextField
                        {...params}
                        label="Enter Cityname"
                        variant="outlined"
                        InputProps={{
                            ...params.InputProps,
                            endAdornment: (
                                <React.Fragment>
                                    {loading ? <CircularProgress color="inherit" size={20} /> : null}
                                    {params.InputProps.endAdornment}
                                </React.Fragment>
                            ),
                        }}
                    />
                )}
            />
            <Map
                {...{
                    id: "zone",
                    type: "fill",
                    source: {
                        type: "geojson",
                        data: cities
                    }
                }}
                style="mapbox://styles/mapbox/streets-v9"
                containerStyle={{
                    height: '100vh',
                    width: '100vw'
                }}
            ></Map>
        </>
    );
}
