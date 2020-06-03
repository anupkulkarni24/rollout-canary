apiVersion: argoproj.io/v1alpha1
kind: AnalysisTemplate
metadata:
  name: success-rate
spec:
  args:
  - name: ingress
  metrics:
  - name: success-rate
    interval: 10s
    successCondition: result[0] > 0.90
    provider:
      prometheus:
        address: http://10.96.203.83:9090
        query: >+
          sum(
            rate(nginx_ingress_controller_requests{ingress="{{args.ingress}}",status!~"[4-5].*"}[5m]))
            /
            sum(rate(nginx_ingress_controller_requests{ingress="{{args.ingress}}"}[5m])
          )
