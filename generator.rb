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

"visit#{klass}(v ast.Visitor, n *ast.#{klass}) interface{} "

# puts <<~EOS
# func visit#{klass}(v ast.Visitor, n *ast.#{klass}) interface{} {
# return nil
# }
# 
# EOS

# ** Node Method
#   puts <<~EOS
# func (n *#{klass}) GetLocation() *Location {
#   return n.Location
# }
#
# EOS
#   puts <<~EOS
# func (n *#{klass}) GetType() string {
#   return "#{klass}"
# }
# EOS


#   puts <<~EOS
# func (v *SoqlChecker) Visit#{klass}(n *ast.#{klass}) interface{} {
# return visit#{klass}(v, n)
# }
#
# EOS
end
