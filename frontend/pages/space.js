import Head from 'next/head';
import Script from 'next/script';
import { useState, useEffect, useRef } from 'react';
import { useRouter } from 'next/router';
import DataContext from '../components/DataContext';
import React, { useContext } from 'react';
import axios from 'axios';
import interact from 'interactjs';

export default function Space() {
  const context = useContext(DataContext);
  const router = useRouter();

  const [message, setMessage] = useState('');
  const [isLoading, setIsLoading] = useState(true);
  const [folders, setFolders] = useState([]);
  const [files, setFiles] = useState([]);
  const [draggedItem, setDraggedItem] = useState(null);
  const containerRef = useRef(null);
  const [isDragOver, setIsDragOver] = useState(false);
  const [uploadedItems, setUploadedItems] = useState([]);

  useEffect(() => {
    getSpace();
  }, []);

  const DesktopItem = ({ id, icon, label, filefolderbool, initialX, initialY }) => {
    const itemRef = useRef(null);
  
    useEffect(() => {
      const item = itemRef.current;
  
      interact(item).draggable({
        listeners: {
          move(event) {
            const target = event.target;
            const x = (parseFloat(target.getAttribute('data-x')) || 0) + event.dx;
            const y = (parseFloat(target.getAttribute('data-y')) || 0) + event.dy;
  
            target.style.transform = `translate(${x}px, ${y}px)`;
            target.setAttribute('data-x', x);
            target.setAttribute('data-y', y);
  
            // Bring the dragged item to the front by setting a higher zIndex
            target.style.zIndex = 1;
          },
          end(event) {
            const target = event.target;
  
            // Set the zIndex back to 0 after the drag ends
            target.style.zIndex = 0;
          },
        },
        modifiers: [
          interact.modifiers.restrictRect({
            restriction: containerRef.current,
          }),
        ],
        inertia: true,
      });
    }, []);
  
    return (
      <div
        ref={itemRef}
        className="draggable dropzone box d-flex justify-center align-items-center flex-column p-2"
        style={{
          flexShrink: 0,
          zIndex: 0, // All items initially have zIndex 0
          transform: `translate(${initialX}px, ${initialY}px)`,
        }}
        data-item={icon}
        data-x={initialX}
        data-y={initialY}
      >
        <svg width="40" height="40" style={{ margin: '5px', flexShrink: 0, zIndex: 1 }} viewBox="0 0 24 24">
          {filefolderbool ? (
            // SVG for file
            <path fill="#37474F" d="M13,9H18.5L13,3.5V9M6,2H14L20,8V20A2,2 0 0,1 18,22H6C4.89,22 4,21.1 4,20V4C4,2.89 4.89,2 6,2M15,18V16H6V18H15M18,14V12H6V14H18Z" />
          ) : (
            // SVG for folder
            <path fill="#37474F" d="M20,18H4V8H20M20,6H12L10,4H4C2.89,4 2,4.89 2,6V18A2,2 0 0,0 4,20H20A2,2 0 0,0 22,18V8C22,6.89 21.1,6 20,6Z" />
          )}
        </svg>
        <p className="text-center display-6 fs-6" style={{ wordWrap: 'break-word', maxWidth: '100%', overflow: 'hidden', textOverflow: 'ellipsis' }}>{label}</p>
      </div>
    );
  };

  const initializeDragAndDrop = () => {
    interact('.draggable').draggable({
      listeners: { move: dragMoveListener },
      inertia: false,
      modifiers: [
        interact.modifiers.restrictRect({
          restriction: 'parent',
          endOnly: true,
        }),
      ],
    });
  };

  const dragMoveListener = (event) => {
    const target = event.target;
    const x = (parseFloat(target.getAttribute('data-x')) || 0) + event.dx;
    const y = (parseFloat(target.getAttribute('data-y')) || 0) + event.dy;

    target.style.transform = `translate(${x}px, ${y}px)`;
    target.setAttribute('data-x', x);
    target.setAttribute('data-y', y);
  };

  const getSpace = async () => {
    try {
      const workspaceId = localStorage.getItem('workspaceId');
      if (workspaceId) {
        context.setWorkspaceId(workspaceId);
      }
      const response = await axios.get(
        `http://localhost:8080/get-directory?folder_id=${workspaceId || context.workspaceId}`
      );
      if (response.status === 200) {
        setIsLoading(false);
        const { folders, files } = response.data;
        setFolders(folders);
        setFiles(files);
      }
    } catch (error) {
      console.log(error);
    }
  };
  const handleDragEnter = (event) => {
    event.preventDefault(); // Prevent the default behavior of dragenter event
    setIsDragOver(true); // Set drag over state to true
  };

  const handleDragLeave = (event) => {
    event.preventDefault(); // Prevent the default behavior of dragleave event
    setIsDragOver(false); // Set drag over state to false
  };

  const handleDrop = async (event) => {
    event.preventDefault(); // Prevent the default behavior of dropping files
    setIsDragOver(false); // Reset drag over state

    const workspaceId = localStorage.getItem('workspaceId');
    if (!workspaceId) {
      console.error('Workspace ID is not available.');
      return;
    }

    const files = event.dataTransfer.files;

    if (files.length === 0) {
      return;
    }

    const formData = new FormData();
    formData.append('file', files[0]); // Assuming you are handling a single file at a time
    formData.append('folder_id', workspaceId);

    try {
      const response = await axios.post('http://localhost:8080/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      if (response.status === 200) {
        // Handle the success response from the backend
        console.log('File uploaded successfully.');
        const newItem = response.data; // Assuming the backend returns the newly created item
        setUploadedItems([...uploadedItems, newItem]);
        console.log(uploadedItems);
      }
    } catch (error) {
      console.error('Error uploading file:', error);
      // Handle the error response from the backend
    }
  };

  useEffect(() => {
    const container = containerRef.current;

    // Add event listeners for dragenter and dragleave to provide visual feedback
    container.addEventListener('dragenter', handleDragEnter);
    container.addEventListener('dragleave', handleDragLeave);

    // Add a drop event listener to the container
    container.addEventListener('drop', handleDrop);

    return () => {
      // Remove the event listeners when the component unmounts
      container.removeEventListener('dragenter', handleDragEnter);
      container.removeEventListener('dragleave', handleDragLeave);
      container.removeEventListener('drop', handleDrop);
    };
  }, []);

  return (
    <div className="container d-flex justify-content-center flex-column my-5">
      <Head>
        <title>Array Component</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
        <link
          href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css"
          rel="stylesheet"
          integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC"
          crossOrigin="anonymous"
        ></link>
      </Head>
      <Script
        src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM"
        crossOrigin="anonymous"
      ></Script>
      <Script src="https://unpkg.com/interactjs/dist/interact.min.js"></Script>
      <style>  {`
    body {
      margin: 30px;
    }
    * {
      font-family: 'Roboto', sans-serif;
    }
    .dropzone {
      position: relative; /* Make sure the child elements are positioned relative to this */
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 30px;
      border-radius: 10px;
    }
    .dropzone:hover {
      transform: scale(1.2);
    }

    /* Part 2 */
    .box {
      width: 110px;
      height: 130px;
      cursor: move;
      user-select: none;
    }
    .draggable {
        position: absolute;
        display: flex;
        flex-direction: column;
        align-items: center;
        text-align: center;
        padding: 10px;
        border-radius: 4px;
        transition: box-shadow 0.2s;
        cursor: pointer;
      }
    
      .draggable:hover {
        background-color: #ffffff;
        border: 1px solid #ccc;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);
        opacity: 0.8;
      }
  `}</style>
      <h1 className="text-primary d-flex justify-content-center">Workspace</h1>
      <div
        ref={containerRef}
        id="file-dropzone"
        className="card container desktop"
        style={{
            width: '1000px', 
            height: '500px',
            position: 'relative',
            overflow: 'hidden', }}
        onDragOver={(e) => e.preventDefault()} // Allow dropping by preventing the default behavior of dragover event
      >
        {folders.map((folder, index) => (
          <DesktopItem
            key={folder.id}
            icon={index + 1}
            label={folder.name}
            filefolderbool={0}
            initialX={index * 120}
            initialY={0}
          />
        ))}
        {files.map((file, index) => (
          <DesktopItem
            key={file.id}
            icon={index + folders.length + 1}
            label={file.name}
            filefolderbool={1}
            initialX={index * 120}
            initialY={0}
          />
        ))}
        {uploadedItems.map((item) => (
      <DesktopItem
        key={item.id}
        icon={item.icon} // Update with the appropriate icon
        label={item.name} // Update with the appropriate label
        filefolderbool={1} // Update with the appropriate value
        initialX={item.x} // Update with the appropriate X-coordinate
        initialY={item.y} // Update with the appropriate Y-coordinate
      />
    ))}
      </div>
    </div>
  );
}
