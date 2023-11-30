import React from 'react';
import './footer.css'

const Footer = ({ aboutTitle, aboutText, contactInfo, phoneNumber, email, authorInfo, authorName, imageSrc }) => {
    const currentYear = new Date().getFullYear();

    return (
        <footer aria-hidden="true">
            <div className="footer-area background-dark footer-padding">
                <div className="container">
                    <div className="col-lg-3 col-md-4 col-sm-6">
                        <div className="footer-section mb-50">
                            <h4 className="footer-heading">{aboutTitle}</h4>
                            <p className="footer-text">{aboutText}</p>
                        </div>
                    </div>
                    <div className="col-lg-3 col-md-4 col-sm-5">
                        <div className="footer-section mb-50">
                            <h4 className="footer-heading">Informacion de contacto:</h4>
                            <ul>
                                <li>Phone: {phoneNumber}</li>
                                <li>Email: {email}</li>
                            </ul>
                        </div>
                    </div>
                    <div className="col-lg-3 col-md-4 col-sm-5">
                        <div className="footer-section">
                            <h4 className="footer-heading">Informacion Adicional</h4>
                            <p className="footer-text">{authorInfo}</p>
                        </div>
                    </div>
                </div>
            </div>
            <div className="footer-bottom background-dark">
                <div className="container">
                    <div className="footer-border">
                        <div className="row justify-content-between align-items-center">
                            <div className="col-lg-10">
                                <div className="footer-copy-right">
                                    <p>&copy; {currentYear} All rights reserved</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </footer>
    );
};

export default Footer;