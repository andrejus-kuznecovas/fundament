import React, { useState, useEffect } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { notesAPI } from '../services/api';
import { Note, CreateNoteData, UpdateNoteData } from '../types';

const NotesPage: React.FC = () => {
    const [notes, setNotes] = useState<Note[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState('');
    const [isCreating, setIsCreating] = useState(false);
    const [editingId, setEditingId] = useState<number | null>(null);
    const [newNoteContent, setNewNoteContent] = useState('');
    const [editContent, setEditContent] = useState('');

    const { user, logout } = useAuth();

    // Load notes on component mount
    useEffect(() => {
        loadNotes();
    }, []);

    const loadNotes = async () => {
        try {
            setIsLoading(true);
            const fetchedNotes = await notesAPI.getAll();
            setNotes(fetchedNotes);
            setError('');
        } catch (err: any) {
            setError('Failed to load notes. Please try again.');
        } finally {
            setIsLoading(false);
        }
    };

    const handleCreateNote = async () => {
        if (!newNoteContent.trim()) return;

        try {
            setIsCreating(true);
            const newNote = await notesAPI.create({ content: newNoteContent.trim() });
            setNotes(prev => [newNote, ...prev]);
            setNewNoteContent('');
            setError('');
        } catch (err: any) {
            setError('Failed to create note. Please try again.');
        } finally {
            setIsCreating(false);
        }
    };

    const handleUpdateNote = async (id: number) => {
        if (!editContent.trim()) return;

        try {
            const updatedNote = await notesAPI.update(id, { content: editContent.trim() });
            setNotes(prev => prev.map(note => note.id === id ? updatedNote : note));
            setEditingId(null);
            setEditContent('');
            setError('');
        } catch (err: any) {
            setError('Failed to update note. Please try again.');
        }
    };

    const handleDeleteNote = async (id: number) => {
        if (!window.confirm('Are you sure you want to delete this note?')) return;

        try {
            await notesAPI.delete(id);
            setNotes(prev => prev.filter(note => note.id !== id));
            setError('');
        } catch (err: any) {
            setError('Failed to delete note. Please try again.');
        }
    };

    const startEditing = (note: Note) => {
        setEditingId(note.id);
        setEditContent(note.content);
    };

    const cancelEditing = () => {
        setEditingId(null);
        setEditContent('');
    };

    if (isLoading) {
        return (
            <div className="min-h-screen flex items-center justify-center">
                <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-gray-50">
            {/* Header */}
            <header className="bg-white shadow">
                <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div className="flex justify-between items-center py-6">
                        <div>
                            <h1 className="text-3xl font-bold text-gray-900">My Notes</h1>
                            <p className="text-gray-600">Welcome back, {user?.email}</p>
                        </div>
                        <button
                            onClick={logout}
                            className="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-md text-sm font-medium"
                        >
                            Logout
                        </button>
                    </div>
                </div>
            </header>

            <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
                <div className="px-4 py-6 sm:px-0">
                    {/* Error message */}
                    {error && (
                        <div className="mb-4 rounded-md bg-red-50 p-4">
                            <div className="text-sm text-red-700">{error}</div>
                        </div>
                    )}

                    {/* Create new note */}
                    <div className="bg-white overflow-hidden shadow rounded-lg mb-6">
                        <div className="p-6">
                            <h3 className="text-lg font-medium text-gray-900 mb-4">Create New Note</h3>
                            <div className="flex space-x-4">
                                <textarea
                                    className="flex-1 border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                    rows={3}
                                    placeholder="Write your note here..."
                                    value={newNoteContent}
                                    onChange={(e) => setNewNoteContent(e.target.value)}
                                    onKeyDown={(e) => {
                                        if (e.key === 'Enter' && !e.shiftKey) {
                                            e.preventDefault();
                                            handleCreateNote();
                                        }
                                    }}
                                />
                                <button
                                    onClick={handleCreateNote}
                                    disabled={!newNoteContent.trim() || isCreating}
                                    className="bg-blue-600 hover:bg-blue-700 disabled:bg-blue-300 text-white px-6 py-2 rounded-md font-medium disabled:cursor-not-allowed"
                                >
                                    {isCreating ? (
                                        <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
                                    ) : (
                                        'Add Note'
                                    )}
                                </button>
                            </div>
                        </div>
                    </div>

                    {/* Notes list */}
                    <div className="space-y-4">
                        {notes.length === 0 ? (
                            <div className="text-center py-12">
                                <p className="text-gray-500 text-lg">No notes yet. Create your first note above!</p>
                            </div>
                        ) : (
                            notes.map((note) => (
                                <div key={note.id} className="bg-white overflow-hidden shadow rounded-lg">
                                    <div className="p-6">
                                        {editingId === note.id ? (
                                            // Edit mode
                                            <div className="space-y-4">
                                                <textarea
                                                    className="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                                    rows={4}
                                                    value={editContent}
                                                    onChange={(e) => setEditContent(e.target.value)}
                                                />
                                                <div className="flex space-x-2">
                                                    <button
                                                        onClick={() => handleUpdateNote(note.id)}
                                                        className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-md text-sm font-medium"
                                                    >
                                                        Save
                                                    </button>
                                                    <button
                                                        onClick={cancelEditing}
                                                        className="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded-md text-sm font-medium"
                                                    >
                                                        Cancel
                                                    </button>
                                                </div>
                                            </div>
                                        ) : (
                                            // View mode
                                            <div>
                                                <div className="flex justify-between items-start mb-4">
                                                    <p className="text-gray-900 whitespace-pre-wrap flex-1">{note.content}</p>
                                                    <div className="flex space-x-2 ml-4">
                                                        <button
                                                            onClick={() => startEditing(note)}
                                                            className="text-blue-600 hover:text-blue-800 text-sm font-medium"
                                                        >
                                                            Edit
                                                        </button>
                                                        <button
                                                            onClick={() => handleDeleteNote(note.id)}
                                                            className="text-red-600 hover:text-red-800 text-sm font-medium"
                                                        >
                                                            Delete
                                                        </button>
                                                    </div>
                                                </div>
                                                <div className="text-sm text-gray-500">
                                                    Created: {new Date(note.created_at).toLocaleDateString()}
                                                    {note.updated_at !== note.created_at && (
                                                        <span className="ml-4">
                                                            Updated: {new Date(note.updated_at).toLocaleDateString()}
                                                        </span>
                                                    )}
                                                </div>
                                            </div>
                                        )}
                                    </div>
                                </div>
                            ))
                        )}
                    </div>
                </div>
            </main>
        </div>
    );
};

export default NotesPage;
