
import { HomeIcon, ClipboardIcon, ChatBubbleLeftIcon } from '@heroicons/react/24/outline';
import Link from 'next/link';
import Image from 'next/image';

export default function Sidebar() {
  return (
    <div className="w-64 bg-white border-r border-gray-200 p-4">
      <div className="flex items-center mb-8">

        <div className="h-8 w-8 bg-blue-600 rounded-lg flex items-center justify-center text-white font-bold">
          T
        </div>
        <span className="ml-2 text-xl font-bold text-gray-900">TrackIt</span>

      </div>
      <nav className="space-y-2">
        <NavItem href="/dashboard" icon={HomeIcon} text="Dashboard" />
        <NavItem href="/tasks" icon={ClipboardIcon} text="Tasks" />
        <NavItem href="/chat" icon={ChatBubbleLeftIcon} text="AI Assistant" />
      </nav>
    </div>
  );
}

function NavItem({ href, icon: Icon, text }: { href: string; icon: any; text: string }) {
  return (
    <Link
      href={href}
      className="flex items-center px-4 py-2 text-gray-600 hover:bg-primary-50 hover:text-primary-600 rounded-lg transition-colors"
    >
      <Icon className="h-5 w-5 mr-3" />
      {text}
    </Link>
  );
}