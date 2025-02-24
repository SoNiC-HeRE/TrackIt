import Link from "next/link"
import { LayoutDashboard, Check, Hourglass, Star, UsersRound, ArrowRightCircle, MenuSquare } from "lucide-react"
import { Button } from "@/components/ui/button"
import { AuroraBackground } from "@/components/ui/aurora-background"

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-white via-gray-50 to-white text-gray-900 flex flex-col">
      <header className="sticky top-0 z-50 w-full border-b border-gray-300 bg-white/90 backdrop-blur supports-[backdrop-filter]:bg-white/60">
        <div className="container flex h-16 items-center justify-between px-4 md:px-8">
          <div className="flex items-center gap-3">
            <LayoutDashboard className="h-7 w-7 text-blue-400" />
            <span className="bg-gradient-to-r from-blue-400 to-violet-400 bg-clip-text text-xl font-bold text-transparent">
              TrackIt
            </span>
          </div>
          
          <Button variant="ghost" className="md:hidden">
            <MenuSquare className="h-5 w-5" />
          </Button>
          
          <nav className="hidden md:flex items-center gap-6">
            <Link href="/login">
              <Button variant="ghost" className="text-gray-700 hover:text-gray-900 hover:bg-gray-200">
                Sign In
              </Button>
            </Link>
            <Link href="/register">
              <Button className="bg-blue-500 hover:bg-blue-600 transition-colors">
                Start for Free
                <Star className="ml-2 h-4 w-4" />
              </Button>
            </Link>
          </nav>
        </div>
      </header>

      <AuroraBackground>
        <main className="flex-1 overflow-hidden">
          <section className="relative overflow-hidden py-20 h-full flex items-center justify-center px-4 md:px-8">
            <div className="container relative z-10 text-center">
              <div className="mx-auto flex max-w-fit items-center gap-3 rounded-full bg-gray-200/50 px-4 py-2 backdrop-blur">
                <span className="flex items-center gap-2 text-sm font-medium text-gray-700">
                  <Star className="h-4 w-4 text-blue-400" />
                  Leading AI Task Management Solution
                </span>
                <div className="flex items-center gap-2 rounded-full bg-blue-500/20 px-2 py-0.5 text-xs text-blue-400">
                  Rated 4.9/5
                </div>
              </div>
              
              <div className="mx-auto mt-8 max-w-4xl space-y-6">
                <h1 className="text-4xl font-bold tracking-tight sm:text-5xl md:text-6xl lg:text-7xl">
                  Boost Your Productivity with {" "}
                  <span className="bg-gradient-to-r from-blue-400 to-violet-400 bg-clip-text text-transparent">
                    AI-Powered
                  </span>{" "}
                  Task Management
                </h1>
                <p className="mx-auto max-w-2xl text-lg text-gray-700 md:text-xl">
                  Simplify your workflow with intelligent task automation. Leverage AI to prioritize, organize, and get more done effortlessly.
                </p>
                
                <div className="mx-auto mt-10 flex max-w-fit flex-col gap-4 sm:flex-row">
                  <Link href="/register">
                    <Button size="lg" className="bg-blue-500 hover:bg-blue-600 transition-colors group">
                      Get Started
                      <ArrowRightCircle className="ml-2 h-4 w-4 transition-transform group-hover:translate-x-1" />
                    </Button>
                  </Link>
                </div>
              </div>
            </div>
            
            <div className="absolute inset-0 -z-10 bg-[radial-gradient(circle_500px_at_50%_200px,#bee3f8,transparent)]" />
          </section>
        </main>
      </AuroraBackground>
      
      <footer className="border-t border-gray-300">
        <div className="container flex h-16 items-center justify-between px-4 md:px-8">
          <div className="flex items-center gap-2">
            <LayoutDashboard className="h-5 w-5 text-blue-400" />
            <span className="bg-gradient-to-r from-blue-400 to-violet-400 bg-clip-text text-sm font-semibold text-transparent">
              TrackIt
            </span>
          </div>
          <p className="text-sm text-gray-600">© Crafted with 💓 by Sriyansh Shivam </p>
        </div>
      </footer>
    </div>
  )
}