import org.graphstream.graph.implementations.SingleGraph
import org.graphstream.graph.{Edge, Node}

object HelloGraphStream {

  val USER_DIR = System.getProperty("user.dir")
  val STYLE = "stylesheet.css"

  def main(args: Array[String]): Unit = {

    println("Hello.java World!")

    // create graph
    val graph = new SingleGraph("hello")
    graph.addAttribute("ui.stylesheet", "url('file://" + USER_DIR + "/" + STYLE + "')")
    graph.addAttribute("ui.antialias")

    // create nodes
    var node: Node = graph.addNode("A")
    node.addAttribute("ui.label", "A")
    node = graph.addNode("B")
    node.addAttribute("ui.label", "B")
    node = graph.addNode("C")
    node.addAttribute("ui.label", "C")

    // create edges
    var edge: Edge = graph.addEdge("AB", "A", "B", true)
    edge = graph.addEdge("BC", "B", "C", true)
    // alternatively
    graph.addEdge[Edge]("CA", "C", "A", true)

    // display graph
    graph.display()
  }

}
