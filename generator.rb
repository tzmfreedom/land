#!/usr/bin/env ruby

f = open('node.txt')
f.each do |line|
  klass = line.strip
#   puts <<~EOS
# func (n *#{klass}) Accept(v Visitor) interface{} {
#   return v.Visit#{klass}(n)
# }
#
# EOS

# puts <<~EOS
# func visit#{klass}(v ast.Visitor, n *ast.#{klass}) interface{} {
# return nil
# }
# 
# EOS

# ** Node Method
#   puts <<~EOS
# func (n *#{klass}) SetParent(parent ast.Node) Node {
#   n.Parent = parent
# }
#
# EOS
#   puts <<~EOS
# func (n *#{klass}) GetType() string {
#   return "#{klass}"
# }
# EOS


  puts <<~EOS
func (v *SoqlChecker) Visit#{klass}(n *ast.#{klass}) interface{} {
return visit#{klass}(v, n)
}

EOS
end
