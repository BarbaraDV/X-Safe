import React from "react";
import './styles.css'




const HeadLight = () => {
    return (
        <div
            key="0"
            className="headlights"
            style={{
                opacity: '0.7846',
                rotate: 'none',
                scale: 'none',
                transform: 'translate3d(0px, -31.7096%, 0px)',
                translate: 'none'
            }}
        >


            <div
                className="headlight-center"
                style={{
                    backgroundImage: 'linear-gradient(312.93deg, rgb(206, 46, 238), rgb(15, 195, 233))'
                }}
            />


            <div
                className="headlights-medium right"
                style={{
                    backgroundImage: 'linear-gradient(312.93deg, rgb(206, 46, 238), rgb(15, 195, 233))'
                }}
            />


            <div
                className="headlights-medium left"
                style={{
                    backgroundImage: 'linear-gradient(312.93deg, rgb(206, 46, 238), rgb(15, 195, 233))'
                }}
            />


            <div
                className="headlights-corner left"
                style={{
                    backgroundImage: 'linear-gradient(312.93deg, rgb(206, 46, 238), rgb(15, 195, 233))'
                }}
            />


            <div
                className="headlights-corner right"
                style={{
                    backgroundImage: 'linear-gradient(312.93deg, rgb(206, 46, 238), rgb(15, 195, 233))'
                }}
            />


        </div>
    )
}


export default HeadLight;