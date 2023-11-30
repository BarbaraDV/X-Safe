import React, { useEffect, useRef, useState } from 'react';
import './sections.css'




const NeonCard = ({ title, description, icon }) => {
    const [isVisible, setIsVisible] = useState(false);
    const containerRef = useRef(null);

    useEffect(() => {
        const observer = new IntersectionObserver(
            ([entry]) => {
                if (entry.isIntersecting) {
                    setIsVisible(true);
                    observer.disconnect();
                }
            },
            { threshold: 0.1 }
        );

        if (containerRef.current) {
            observer.observe(containerRef.current);
        }

        return () => observer.disconnect();
    }, []);


    return (
        <div className={`cards-container ${isVisible ? 'visible' : ''}`} ref={containerRef}>
            <div className={`card `}>
                {icon}
                <h2 className='card-title x-safe-text'>{title}</h2>
                <p className='card-text'>{description}</p>
            </div>
        </div>
    );
};


const ObjectivesSection = () => {
    return (
        <section className="objectives-section" >
            <NeonCard title="Misión" icon={<svg
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
                xmlnsXlink="http://www.w3.org/1999/xlink"
                viewBox="0 0 32 32"
                xmlSpace="preserve"
                width="64px"
                height="64px"
                fill="white"
                stroke="white"
            >
                <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                <g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g>
                <g id="SVGRepo_iconCarrier">
                    <style type="text/css">
                        {`
            .st0 {
              fill: none;
              stroke: #fffff;
              stroke-width: 2;
              stroke-linecap: round;
              stroke-linejoin: round;
              stroke-miterlimit: 10;
            }
          `}
                    </style>
                    <line className="st0" x1="16" y1="16" x2="22" y2="10" />
                    <polygon className="st0" points="30,6 26,6 26,2 22,6 22,10 26,10 " />
                    <circle className="st0" cx="16" cy="16" r="6" />
                    <path
                        className="st0"
                        d="M27,9c1.3,2,2,4.4,2,7c0,7.2-5.8,13-13,13S3,23.2,3,16S8.8,3,16,3c2.6,0,5,0.7,7,2"
                    />
                </g>
            </svg>} description="Nuestra misión es proporcionar una protección inquebrantable en Twitter, identificando todo tipo de vulnerabilidades. Nos esforzamos por ofrecer actualizaciones que anticipen y neutralicen las amenazas antes de que impacten a los usuarios." />
            <NeonCard title="Visión"
                icon={<svg
                    width="64px"
                    height="64px"
                    viewBox="0 0 24 24"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    stroke="white"
                >
                    <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                    <g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g>
                    <g id="SVGRepo_iconCarrier">
                        <path
                            d="M9 4.45962C9.91153 4.16968 10.9104 4 12 4C16.1819 4 19.028 6.49956 20.7251 8.70433C21.575 9.80853 22 10.3606 22 12C22 13.6394 21.575 14.1915 20.7251 15.2957C19.028 17.5004 16.1819 20 12 20C7.81811 20 4.97196 17.5004 3.27489 15.2957C2.42496 14.1915 2 13.6394 2 12C2 10.3606 2.42496 9.80853 3.27489 8.70433C3.75612 8.07914 4.32973 7.43025 5 6.82137"
                            stroke="white"
                            strokeWidth="1.5"
                            strokeLinecap="round"
                        />
                        <path
                            d="M15 12C15 13.6569 13.6569 15 12 15C10.3431 15 9 13.6569 9 12C9 10.3431 10.3431 9 12 9C13.6569 9 15 10.3431 15 12Z"
                            stroke="white"
                            strokeWidth="1.5"
                        />
                    </g>
                </svg>}
                description="Imaginamos un futuro donde Twitter sea sinónimo de seguridad y fiabilidad. Aspiramos a construir una comunidad donde los usuarios se sientan completamente seguros, empoderados por nuestra plataforma para interactuar y compartir sin temor." />
            <NeonCard title="Objetivo"
                icon={<svg
                    fill="#ffffff"
                    height="64px"
                    width="64px"
                    viewBox="0 0 329.996 329.996"
                    xmlns="http://www.w3.org/2000/svg"
                    xmlnsXlink="http://www.w3.org/1999/xlink"
                    xmlSpace="preserve"
                    stroke="#ffffff"

                >
                    <path d="M264.207,129.996l52.504-65.629c3.603-4.503,4.305-10.671,1.807-15.869c-2.498-5.197-7.754-8.502-13.52-8.502l-125,0.004V15 c0-3.979-1.58-7.794-4.394-10.607C172.791,1.58,168.975,0,164.997,0l-140,0.005c-8.283,0-14.999,6.716-14.999,15v0.009v149.982v150 c0,8.284,6.716,15,15,15s15-6.716,15-15v-135h110v25c0,8.284,6.716,15,15,15h140c5.766,0,11.022-3.305,13.52-8.502 c2.498-5.197,1.796-11.366-1.807-15.869L264.207,129.996z M39.998,149.996V30.004l110-0.004v25v94.996H39.998z M179.998,189.996v-25 V69.999l93.79-0.002l-40.503,50.628c-4.383,5.479-4.383,13.263,0,18.741l40.504,50.629H179.998z" />
                </svg>}
                description="Evolución constante en la detección de vulnerabilidades en Twitter. Nos dedicamos a identificar proactivamente y abordar de manera efectiva las amenazas emergentes en la plataforma, asegurando que nuestros usuarios estén siempre un paso adelante en la seguridad digital." />
        </section>
    );
};

export default ObjectivesSection;