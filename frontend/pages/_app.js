import { useState, useEffect } from 'react';
import DataContext from '../components/DataContext';
import { useRouter } from 'next/router';

function MyApp({ Component, pageProps }) {
    
    const [ appKey ] = useState(process.env.APP_KEY);
    const [ userId, setUserId ] = useState('');
    const [ workspaceId, setWorkspaceId ] = useState('');


    const router = useRouter();    

    console.log("userId: " + userId);

    return (
        <DataContext.Provider value={{ appKey, userId, setUserId, workspaceId, setWorkspaceId }}>
            <Component {...pageProps} />
        </DataContext.Provider>
    );
}

export default MyApp;