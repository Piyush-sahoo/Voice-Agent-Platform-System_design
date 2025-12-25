export default function Home() {
    return (
        <main className="flex min-h-screen flex-col items-center justify-center p-24">
            <div className="text-center">
                <h1 className="text-4xl font-bold mb-4">
                    VAP Platform
                </h1>
                <p className="text-xl text-gray-600 mb-8">
                    Voice AI Platform Dashboard
                </p>

                <div className="grid grid-cols-2 gap-4 max-w-2xl">
                    <a href="/login" className="p-4 border rounded-lg hover:bg-gray-50">
                        <h2 className="text-lg font-semibold">Login →</h2>
                        <p className="text-gray-500">Sign in to your account</p>
                    </a>

                    <a href="/register" className="p-4 border rounded-lg hover:bg-gray-50">
                        <h2 className="text-lg font-semibold">Register →</h2>
                        <p className="text-gray-500">Create a new account</p>
                    </a>

                    <a href="/dashboard" className="p-4 border rounded-lg hover:bg-gray-50">
                        <h2 className="text-lg font-semibold">Dashboard →</h2>
                        <p className="text-gray-500">View your overview</p>
                    </a>

                    <a href="/agents" className="p-4 border rounded-lg hover:bg-gray-50">
                        <h2 className="text-lg font-semibold">Agents →</h2>
                        <p className="text-gray-500">Manage voice agents</p>
                    </a>
                </div>
            </div>
        </main>
    )
}
