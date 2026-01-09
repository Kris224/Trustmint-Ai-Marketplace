// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from 'firebase/auth';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyCqcgPDEPQfaSwbyWGuPvquCkaGx5s-lJs",
  authDomain: "trustm-19ece.firebaseapp.com",
  projectId: "trustm-19ece",
  storageBucket: "trustm-19ece.firebasestorage.app",
  messagingSenderId: "322823059855",
  appId: "1:322823059855:web:f75cc952239d9479adf3fa"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

export { auth };