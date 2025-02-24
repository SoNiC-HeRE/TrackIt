import Link from "next/link"
import { Bot, CheckCircle, Clock, Sparkles, Users } from "lucide-react"

import { Button } from "@/components/ui/button"

export default function Home() {
  return (
    <div className="flex min-h-screen flex-col">
      <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
        <div className="container flex h-14 items-center justify-between">
          <div className="flex items-center gap-2">
            <Bot className="h-6 w-6 text-blue-600" />
            <span className="text-xl font-bold">TrackIt</span>
          </div>
          <nav className="hidden md:flex gap-6">
            <Link className="text-sm font-medium hover:text-blue-600" href="#">
              Features
            </Link>
            <Link className="text-sm font-medium hover:text-blue-600" href="#">
              Solutions
            </Link>
            <Link className="text-sm font-medium hover:text-blue-600" href="#">
              Pricing
            </Link>
            <Link className="text-sm font-medium hover:text-blue-600" href="#">
              About
            </Link>
          </nav>
          <div className="flex items-center gap-4">
            <Link href="/login">
              <Button variant="ghost">Login</Button>
            </Link>
            <Link href="/register">
              <Button className="bg-blue-600 hover:bg-blue-700">
                Try for free
                <Sparkles className="ml-2 h-4 w-4" />
              </Button>
            </Link>
          </div>
        </div>
      </header>

      <main className="flex-1">
        <section className="container space-y-8 py-12 text-center md:py-24 lg:py-32">
          <div className="mx-auto flex max-w-fit items-center gap-2 rounded-full bg-muted px-4 py-1.5">
            <span className="flex items-center gap-1 text-sm font-medium">
              <Sparkles className="h-4 w-4 text-blue-600" />
              #1 AI Task Management Platform
            </span>
            <div className="flex items-center gap-1 rounded-full bg-blue-600 px-2 py-0.5 text-xs text-white">
              4.9/5 Rating
            </div>
          </div>
          <div className="mx-auto max-w-3xl space-y-4">
            <h1 className="text-4xl font-bold tracking-tight sm:text-5xl md:text-6xl lg:text-7xl">
              Supercharge your productivity with{" "}
              <span className="bg-gradient-to-r from-blue-600 to-violet-600 bg-clip-text text-transparent">
                AI-powered
              </span>{" "}
              tasks
            </h1>
            <p className="mx-auto max-w-[700px] text-muted-foreground md:text-xl">
              Streamline your workflow with intelligent task management. Let AI help you prioritize, organize, and
              accomplish more in less time.
            </p>
          </div>
          <div className="mx-auto flex max-w-fit flex-col gap-4 sm:flex-row">
            <Link href="/register">
              <Button size="lg" className="bg-blue-600 hover:bg-blue-700">
                Get started
                <Sparkles className="ml-2 h-4 w-4" />
              </Button>
            </Link>
            <Button size="lg" variant="outline">
              Book a demo
            </Button>
          </div>
        </section>

        <section className="container py-12 md:py-24 lg:py-32">
          <div className="mx-auto grid max-w-5xl gap-12 sm:grid-cols-2 md:grid-cols-4">
            <div className="flex flex-col items-center gap-2 text-center">
              <Bot className="h-12 w-12 text-blue-600" />
              <h3 className="text-xl font-semibold">AI-Powered</h3>
              <p className="text-sm text-muted-foreground">Smart task suggestions and intelligent prioritization</p>
            </div>
            <div className="flex flex-col items-center gap-2 text-center">
              <Clock className="h-12 w-12 text-blue-600" />
              <h3 className="text-xl font-semibold">Real-time</h3>
              <p className="text-sm text-muted-foreground">Instant updates and live collaboration features</p>
            </div>
            <div className="flex flex-col items-center gap-2 text-center">
              <Users className="h-12 w-12 text-blue-600" />
              <h3 className="text-xl font-semibold">Team-friendly</h3>
              <p className="text-sm text-muted-foreground">Seamless task assignment and team management</p>
            </div>
            <div className="flex flex-col items-center gap-2 text-center">
              <CheckCircle className="h-12 w-12 text-blue-600" />
              <h3 className="text-xl font-semibold">Secure</h3>
              <p className="text-sm text-muted-foreground">JWT authentication and data encryption</p>
            </div>
          </div>
        </section>

        <section className="border-t bg-muted/50">
          <div className="container py-12 md:py-24 lg:py-32">
            <div className="mx-auto grid max-w-5xl items-center gap-12 lg:grid-cols-2">
              <div className="flex flex-col gap-4">
                <h2 className="text-3xl font-bold tracking-tight sm:text-4xl md:text-5xl">
                  Let AI handle your task management
                </h2>
                <p className="text-muted-foreground md:text-lg">
                  Our AI-powered system learns from your work patterns to suggest optimal task organization, deadlines,
                  and team assignments. Focus on execution while we handle the planning.
                </p>
                <div className="flex gap-4">
                  <Link href="/register">
                    <Button className="bg-blue-600 hover:bg-blue-700">
                      Start organizing
                      <Sparkles className="ml-2 h-4 w-4" />
                    </Button>
                  </Link>
                </div>
              </div>
              <div className="rounded-xl border bg-background p-8">
                <div className="flex flex-col gap-4">
                  <div className="flex items-center gap-4">
                    <Bot className="h-8 w-8 text-blue-600" />
                    <div className="flex-1">
                      <h4 className="font-semibold">AI Assistant</h4>
                      <p className="text-sm text-muted-foreground">Here are your suggested tasks for today:</p>
                    </div>
                  </div>
                  <div className="space-y-2">
                    <div className="rounded-lg border bg-muted/50 p-3">
                      <p className="text-sm">⏰ Priority: Review Q1 project milestones (2 hours)</p>
                    </div>
                    <div className="rounded-lg border bg-muted/50 p-3">
                      <p className="text-sm">👥 Team: Schedule weekly sync with design team</p>
                    </div>
                    <div className="rounded-lg border bg-muted/50 p-3">
                      <p className="text-sm">📊 Analysis: Prepare monthly performance report</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section className="border-t">
          <div className="container py-12 md:py-24 lg:py-32">
            <div className="mx-auto max-w-3xl space-y-8 text-center">
              <h2 className="text-3xl font-bold tracking-tight sm:text-4xl md:text-5xl">
                Ready to transform your task management?
              </h2>
              <p className="text-muted-foreground md:text-xl">
                Join thousands of teams already using TrackIt to streamline their workflow and boost productivity.
              </p>
              <div className="flex justify-center gap-4">
                <Link href="/register">
                  <Button size="lg" className="bg-blue-600 hover:bg-blue-700">
                    Get started for free
                    <Sparkles className="ml-2 h-4 w-4" />
                  </Button>
                </Link>
                <Button size="lg" variant="outline">
                  View pricing
                </Button>
              </div>
            </div>
          </div>
        </section>
      </main>

      <footer className="border-t">
        <div className="container flex h-14 items-center justify-between">
          <div className="flex items-center gap-2">
            <Bot className="h-5 w-5 text-blue-600" />
            <span className="text-sm font-semibold">TrackIt</span>
          </div>
          <p className="text-sm text-muted-foreground">© {new Date().getFullYear()} TrackIt. All rights reserved.</p>
        </div>
      </footer>
    </div>
  )
}

